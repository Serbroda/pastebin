import React, {FC, useEffect, useState} from 'react'
import {PasteDto} from "../models/PasteDto";
import {pasteService} from "../services/Services";
import Container from "../components/Container";

const Home: FC = () => {
    const [paste, setPaste] = useState<PasteDto | undefined>(undefined);

    useEffect(() => {
        pasteService.createPaste({
            content: 'Test React'
        }).then(setPaste);
    }, []);

    return (
        <Container>
            <p>{paste?.id}</p>
        </Container>
    )
}

export default Home