![tic-tac-go logo](tictacgo.svg)

__tic-tac-go is tic-tac-toe written in Go__

__Issues:__ Currently doesn't handle tied games well

__Usage:__
To use run *go run tictacgo.go*

Player X will be asked to enter a row for which they type either 0, 1 or 2
Player X will be asked to enter a col for which they type either 0, 1 or 2
An X is then placed in the (row,col) position of the grid

The same happens for Player O

-(0,0)-|-(0,1)-|-(0,2)-
-(1,0)-|-(1,1)-|-(1,2)-
-(2,0)-|-(2,1)-|-(2,2)-

When a Player gets 3 in a row they will be notified that they have won.
