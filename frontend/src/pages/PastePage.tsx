import React, {FC, useEffect, useState} from 'react'
import {PasteDto} from "../models/PasteDto";
import {pasteService} from "../services/Services";
import Container from "../components/Container"

const Paste: FC = () => {
    const [pastes, setPastes] = useState<PasteDto[]>([]);

    useEffect(() => {
        pasteService.getPastes().then(setPastes);
    }, []);

    return (
        <Container>
            <h1 className="text-3xl font-bold underline text-red-600">
                Session Pastes
            </h1>

            {pastes.map((p) => <p>{p.id}</p>)}
        </Container>
    )
}

export default Paste