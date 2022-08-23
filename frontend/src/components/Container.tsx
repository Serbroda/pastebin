import React from 'react'
import {Link} from "react-router-dom";

const Container: React.FC<{children: React.ReactNode}> = ({children}) => {
    return (
        <div>
            <nav
                style={{
                    borderBottom: "solid 1px",
                    paddingBottom: "1rem",
                }}
            >
                <Link to="/">Home</Link> |{" "}
                <Link to="/pastes">Pastes</Link>
            </nav>

            {children}
        </div>
    )
}

export default Container