{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Use camelCase" #-}
{-# HLINT ignore "Use null" #-}
{-# HLINT ignore "Eta reduce" #-}
{-# HLINT ignore "Avoid lambda" #-}
module Ejercicio1
    (main) where

--Encuentro  si se encuentra el primer char y coincide con letra del texto luego el segundo y en esa dinamica sigue
sub_cadenasAux :: [Char] -> [Char] -> Bool
sub_cadenasAux a b 
    | a == [] = True
    | b == [] = False
    | head a  == head b = sub_cadenasAux (tail a) (tail b)
    | otherwise = sub_cadenasAux a (tail b)

sub_cadenas :: [Char] -> [String] -> [String]
sub_cadenas a b =
    filter (\x -> sub_cadenasAux a x)b
main :: IO()
main = do
    print(sub_cadenas "la" ["la casa", "el perro", "pintando la cerca","la manzana"])
-- Lo que imprime ["la casa","pintando la cerca","la casa","la manzana"]