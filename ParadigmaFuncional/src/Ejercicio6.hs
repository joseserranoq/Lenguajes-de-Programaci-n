module Ejercicio6
    ( main
    ) where

--miembro 
miembro elem lista =
    any comparacion lista
    --any (==True) (map comparacion lista)
    where
        comparacion x = x == elem

--remove_if
remove_if fun lista =
    concatMap (\x -> if (fun x) then []
                     else [x] )
                lista

vecinos nodo grafo = do
    let resultado =  concatMap (\x -> if (head (head x)) == nodo then [x]
                          else [] )
                            grafo
    if resultado == [] then []
    else head (tail (head resultado))

extender ruta grafo =
    remove_if (==[])
    (map (\x -> if miembro x ruta then []
                       else x:ruta )
        (vecinos (head ruta) grafo))

prof ini fin = prof_aux fin [[ini]]

prof_aux fin rutas grafo
  | rutas==[] = []
  | (fin == (head (head rutas))) =
reverse (head rutas)
  | otherwise =
prof_aux fin ((extender (head rutas) grafo)++(tail rutas)) grafo

main :: IO ()
main = do
    let grafo=[ [["i"],["a","b"]],      --        a ---- c ---- x
                [["a"],["i","c","d"]],  --      /   \  /
                [["b"],["i","c","d"]],  --     /     \/
                [["c"],["a","b","x"]],  --   i       X
                [["d"],["a","b","f"]],  --     \    / \
                [["f"],["d"]],          --      \  /   \
                [["x"],["c"]]           --        b ---- d ---- f
                 ]

    print(prof "i" "f" grafo)