import React, { useEffect, useState } from 'react';
import { HarmoVisLayers, Container, BasedProps, BasedState, connectToHarmowareVis, MovesLayer, Movesbase, MovesbaseOperation, DepotsLayer, DepotsData, LineMapLayer, LineMapData } from 'harmoware-vis';
import io from "socket.io-client";

const MAPBOX_TOKEN = 'pk.eyJ1IjoicnVpaGlyYW5vIiwiYSI6ImNqdmc0bXJ0dTAzZDYzem5vMmk0ejQ0engifQ.3k045idIb4JNvawjppzqZA'
const socket: SocketIOClient.Socket = io();

class Harmoware extends Container<BasedProps & BasedState> {
    render() {
        return (<HarmowarePage {...this.props} />)
    }
}

const HarmowarePage: React.FC<BasedProps & BasedState> = (props) => {
    const { actions, depotsData, viewport, movesbase, movedData, routePaths, clickedObject } = props

    const [movesdata, setMovesdata] = useState<Movesbase[]>([])

    const setAgents = (data: any) => {
        const time = Date.now() / 1000; // set time as now. (If data have time, ..)
        const newMovesbase: Movesbase[] = [];
        // useEffect内では外側のstateは初期化時のままなので、set関数内で過去のstateを取得する必要がある
        setMovesdata((movesdata) => {
            //console.log("socketData: ", movesdata);
            movesdata.forEach((movedata) => {
                let isExist = false;
                let color = [0, 200, 120];
                data.forEach((value: any) => {
                    const { type, id, latitude, longitude } = JSON.parse(
                        value
                    );
                    //console.log("id, type: ", id, movedata.type)
                    if (id === movedata.type) {
                        //console.log("match")
                        // 存在する場合、更新
                        newMovesbase.push({
                            ...movedata,
                            operation: [
                                ...movedata.operation,
                                {
                                    elapsedtime: time,
                                    position: [longitude, latitude, 0],
                                    color
                                }
                            ]
                        });
                        isExist = true
                    }
                    if (!isExist) {
                        // 存在しない場合、新規作成
                        let color = [0, 255, 0];
                        newMovesbase.push({
                            type: id,
                            operation: [
                                {
                                    elapsedtime: time,
                                    position: [longitude, latitude, 0],
                                    color
                                }
                            ]
                        });
                    }

                })
            })


            return newMovesbase
        })
        actions.updateMovesBase(newMovesbase);
    }

    useEffect(() => {
        socket.on("agents", (data: any) => setAgents(data));


        console.log(process.env);
        if (actions) {
            actions.setViewport({
                ...props.viewport,
                longitude: 136.9831702,
                latitude: 35.1562909,
                width: window.screen.width,
                height: window.screen.height,
                zoom: 16
            })
            actions.setSecPerHour(3600);
            actions.setLeading(2)
            actions.setTrailing(5)
        }
    }, [])

    return (
        <div className="App">
            <HarmoVisLayers
                viewport={viewport} actions={actions}
                mapboxApiAccessToken={MAPBOX_TOKEN}
                layers={[

                ]}
            />
        </div>
    );
}

export default connectToHarmowareVis(Harmoware);