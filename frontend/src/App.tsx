import React from 'react';
import {BrowserRouter, Route, Routes} from "react-router-dom";
import Home from "./pages/HomePage";
import Paste from "./pages/PastePage";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home/>}/>
                <Route path="/pastes" element={<Paste/>}/>
            </Routes>
        </BrowserRouter>
    );
}

export default App;
