{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Use camelCase" #-}
module Ejercicio2
    (main) where
--el tercer parametro es para comenzar la busqueda con la lista entera nuevamente
sub_conjunto :: Eq k => [k] -> [k] -> Int -> Bool
sub_conjunto a b c
    | a == [] = True
    | c == length b = False
    | b!!c == head a = sub_conjunto (tail a) b 0
    | otherwise = sub_conjunto a b (c+1) 

main :: IO()
main = do
    print(sub_conjunto [] [1,2,3,4,5] 0)
    print(sub_conjunto [1,2,3] [1,2,3,4,5] 0)
    print(sub_conjunto [1,2,6] [1,2,3,4,5] 0)
    print(sub_conjunto [14,5,235] [5,734,96,14,976,235] 0)
    {-
    Devuelve respectivamente:
    True
    True
    False
    True 
    -}