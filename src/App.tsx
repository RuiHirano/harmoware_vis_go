import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Provider } from 'react-redux';
import { createStore } from 'redux';
import { getCombinedReducer } from 'harmoware-vis';
import { Harmoware } from './views'



const store = createStore(getCombinedReducer());

function App() {
  return (

    <Provider store={store}>
      <Harmoware />
    </Provider>
  );
}

export default App;
