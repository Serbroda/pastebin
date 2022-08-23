import React from 'react';
import {HashRouter, Route, Routes} from "react-router-dom";
import Home from "./pages/HomePage";
import Paste from "./pages/PastePage";

function App() {
    return (
        <HashRouter>
            <Routes>
                <Route path="/" element={<Home/>}/>
                <Route path="/pastes" element={<Paste/>}/>
            </Routes>
        </HashRouter>
    );
}

export default App;
