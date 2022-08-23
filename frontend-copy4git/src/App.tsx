import React, {useEffect, useState} from 'react';
import {pasteService} from "./services/Services";
import {PasteDto} from "./models/PasteDto";

function App() {
    const [paste, setPaste] = useState<PasteDto | undefined>(undefined);
    const [pastes, setPastes] = useState<PasteDto[]>([]);

    useEffect(() => {
        pasteService.getPastes().then(setPastes);

        pasteService.createPaste({
            content: 'Test React'
        }).then(p => {
            setPaste(p);
            pasteService.getPastes().then(setPastes);
        });
    }, []);

    return (
        <div className="App">
            <h1 className="text-3xl font-bold underline text-red-600">
                Simple React Typescript Tailwind Sample
            </h1>

            <p>{paste?.id}</p>
            <br/>
            <p>Session Pastes</p>
            {pastes.map((p) => <p>{p.id}</p>)}
        </div>
    );
}

export default App;
