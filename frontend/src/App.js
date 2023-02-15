import { createGraphiQLFetcher } from '@graphiql/toolkit';
import { GraphiQL } from 'graphiql';
import React, {useEffect, useState} from 'react';
import './App.css';
import 'graphiql/graphiql.css';



const App = () => {
    let [heroes, setHeroes] = useState(["hasan", "ocak"])
    let [color, setColor] = useState(false)

    let query = `{
            heroes
            {
              name
            }
          }`

    useEffect(() => {
        const beCall = async () => {
            try {
                let res = await  fetch('http://localhost:8085/query', {
                    method: 'POST',
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({
                        query: query
                    })
                })

                let jsonData= await res.json()
                setHeroes(jsonData.data.heroes)
            } catch (err) {
                console.log(err)
            }
        }
        beCall()
    }, [])

    function changeColor(e) {
        e.preventDefault();
        console.log("clicked");
        let colorCopy = color;
        setColor(!colorCopy);
    }

    let backGroundColor =  color? "white": "cyan";

    return (
        <div className="wireframe" style={{backgroundColor: backGroundColor}}>
            <button className="button" onClick={(e) => changeColor(e)}>Change color</button>
            <div className="heroes">
                {heroes.map((hero, idx) => <div id={idx} className="hero">{hero.name}</div>)}
            </div>
            <GraphiQL
                fetcher={createGraphiQLFetcher({
                    url: 'http://localhost:8085/query',
                })}
            />
        </div>
    )
}

export default App;
