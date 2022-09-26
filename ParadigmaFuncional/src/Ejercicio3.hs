module Ejercicio3
    (main) where

aplanar :: Eq lista => [[lista]] -> [lista] -> Int -> [lista]
aplanar a b c
    | c == length a = b
    | a!!c /= []  = aplanar a (b ++ a!!c) (c+1)
    | otherwise = aplanar a b (c+1)
    


main :: IO()
main = do
    let a = [[5,3,1],[1,2],[5,3,14,6,1],[99,44,22]]
    print(aplanar a [] 0)
    -- Lo que imprime
    -- [5,3,1,1,2,5,3,14,6,1,99,44,22]

