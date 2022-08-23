import React, {FC, useEffect, useState} from 'react'
import {PasteDto} from "../models/PasteDto";
import {pasteService} from "../services/Services";
import {Link} from "react-router-dom";

const Paste: FC = () => {
    const [pastes, setPastes] = useState<PasteDto[]>([]);

    useEffect(() => {
        pasteService.getPastes().then(setPastes);
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
                Session Pastes
            </h1>

            {pastes.map((p) => <p>{p.id}</p>)}
        </div>
    )
}

export default Paste