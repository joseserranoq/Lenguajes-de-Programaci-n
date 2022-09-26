{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Redundant bracket" #-}
module Ejercicio5
    ( main) where

data Producto = Producto { nombre :: String
                    , cantidad :: Int
                    , precio :: Int
                     } deriving Show


listarExistenciaMinima :: Int -> [Producto] -> [Producto]
listarExistenciaMinima minimo listaProductos =
    filter minimos listaProductos 
    where
        minimos prod = ((cantidad prod) <= minimo) 

--calcularMontoVenta nombreProducto cantidad
--CALCULAR EL MONTO A PAGAR POR LA CANTIDAD DE PRODUCTOS INDICADA MÁS EL 13% DE IMPUESTO


calcular :: Int -> Int -> Int
calcular a b = a * b

calcularMontoVenta :: Int -> [Producto] -> Int
calcularMontoVenta a b 
-- * 13 `div`100
    | length b == 0 = a * 13 `div`100 + a
    | otherwise = calcularMontoVenta (a + calcular (cantidad (head b)) (precio (head b))) (tail b)
    
main :: IO ()
main = do
    let listaProductos=[
            (Producto "arroz" 15 2500),
            (Producto "frijoles" 4 2000),
            (Producto "leche" 8 1200),
            (Producto "cafe" 12 4500)]
    print(calcularMontoVenta 0 listaProductos)
    -- 123283
