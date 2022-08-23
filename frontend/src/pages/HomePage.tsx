import React, {FC, useEffect, useState} from 'react'
import {PasteDto} from "../models/PasteDto";
import {pasteService} from "../services/Services";
import {Link} from "react-router-dom";

const Home: FC = () => {
    const [paste, setPaste] = useState<PasteDto | undefined>(undefined);

    useEffect(() => {
        pasteService.createPaste({
            content: 'Test React'
        }).then(setPaste);
    }, []);

    return (
        <div>
            <nav
                style={{
                    borderBottom: "solid 1px",
                    paddingBottom: "1rem",
                }}
            >
                <Link to="/">Invoices</Link> |{" "}
                <Link to="/pastes">Expenses</Link>
            </nav>

            <h1 className="text-3xl font-bold underline text-red-600">
                Simple React Typescript Tailwind Sample
            </h1>

            <p>{paste?.id}</p>
        </div>
    )
}

export default Home