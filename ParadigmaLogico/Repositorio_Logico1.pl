%listing busca alguna funcion con el valor at[omico

%Ejercicio 1 del repo, realizar en papel
miembro(N,[N|_]). %Hecho
miembro(N,[_|Tail]) :- miembro(N,Tail). %Regla
% si n en el hecho es false -> quita el primer elemento de la lista mediante lo escrito en regla
% asi sucesivamente hasta encontrar el N o terminar la lista y dar false.
% ejemplo
% miembro(2,[1,2,3])
% 1 == 2 false y miembro([2,3])
% 2  == 2 true.


%Ejercicio 2 del repo
factorial(0,1).
factorial(A,B) :-
		A > 0,
		C is A-1,
		factorial(C,D),
		B is A*D.
% Si el factorial es 0 entonces devuelte true
% Sino entonces pasa por las reglas en donde se saca como resultado B
%ejemplo
%factorial(3,X).
% se reduce A-1 hasta que A llegue a 0 guarda las funciones en la pila
% con sus respectivos valores
% prolog intuye que el valor de B es en el ques se va a dar el resultad
% entones B = 1 * 2 -> B = 2 * 3, termina B=6

%Ejercicio 3
insertar(E,[],[E]).
insertar(E,[A|B],[E|[A|B]]):-  E < A,!.
insertar(E,[A|B],[A|R]):- insertar(E,B,R).
% Se introduce el elemento que a va a ser introducido en la lista del
% segundo parametro es la lista en la que se va a agregar el elemento E
% el tercer parametro evalua si la nueva lista que se genera al insertar
% el elemento.
% ejemplo
% insertar(3,[1,2,4],X).
% busca hasta que E sea menor que la cabeza de la lista para insertarlo
% en la lista en esa posicion.
% ejemplo
% insertar(3,[1,2,4],X).
% 3 < 1 = false, 3< 2 =false, 3<4 = true, se inserta en esa posicion.
% y despliega X=[1,2,3,4]

%Funcion para borrar los 3
deleteall([],X,[]).
deleteall([H|T],X,Result) :-
    H==X,
    deleteall(T,X,Result).
deleteall([H|T],X,[H|Result]) :- deleteall(T,X,Result).
