package main

import (
    "fmt"
    "math/rand"
    "time"
)

func defineSet() int {
    winner_player_one := rand.Intn(1)
    fmt.Println("quien gano?", winner_player_one)
    return winner_player_one
}

func main() {
    c1 := make(chan int)
    c2 := make(chan int)

    var player_one_saca = 1
    var random_player_one int
    var random_player_two int

    var puntos_player_one = 0
    var puntos_player_two = 0

    var set_ganados_player_one = 0
    var set_ganados_player_two = 0

    for {

        if player_one_saca == 1 {
            random_player_one = (rand.Intn(4) + 1)
            random_player_two = (rand.Intn(3) + 1)
        } else {
            random_player_one = (rand.Intn(3) + 1)
            random_player_two = (rand.Intn(4) + 1)
        }

        go func() { time.Sleep(time.Second * 1); c1 <- random_player_one }()
        go func() { time.Sleep(time.Second * 1); c2 <- random_player_two }()

        msg1 := <-c1
        msg2 := <-c2
        if msg1 > 0 && msg2 > 0 {

            if player_one_saca == 1 {
                if msg1 == 3 && (msg2 == 1 || msg2 == 2) {
                    puntos_player_one += 1
                } else if msg1 == 1 || msg1 == 2 {
                    puntos_player_two += 1
                } else if msg1 == 4 {
                    puntos_player_one += 1
                } else if msg1 == 3 && msg2 == 3 {

                    winner_player_one := defineSet()
                    if winner_player_one == 1 {
                        puntos_player_one += 1
                    } else {
                        puntos_player_two += 1
                    }
                } else {
                    puntos_player_two += 1
                }
            } else {
                if msg2 == 3 && (msg1 == 1 || msg1 == 2) {
                    puntos_player_two += 1
                } else if msg2 == 1 || msg2 == 2 {
                    puntos_player_one += 1
                } else if msg2 == 4 {
                    puntos_player_two += 1
                } else if msg1 == 3 && msg2 == 3 {
                    winner_player_one := defineSet()
                    if winner_player_one == 1 {
                        puntos_player_one += 1
                    } else {
                        puntos_player_two += 1
                    }
                } else {
                    puntos_player_one += 1
                }
            }

            if (puntos_player_one >= 4 && ((puntos_player_one - puntos_player_two) > 1)) ||
                (puntos_player_two >= 4 && ((puntos_player_two - puntos_player_one) > 1)) {

                if puntos_player_one >= 4 && ((puntos_player_one - puntos_player_two) > 1) {
                    set_ganados_player_one += 1
                }

                if puntos_player_two >= 4 && ((puntos_player_two - puntos_player_one) > 1) {
                    set_ganados_player_two += 1
                }

                if player_one_saca == 0 {
                    player_one_saca = 1
                    fmt.Println("cambio de saque saca player one")
                    puntos_player_one = 0
                    puntos_player_two = 0
                } else {
                    player_one_saca = 0
                    fmt.Println("cambio de saque saca player two")
                    puntos_player_one = 0
                    puntos_player_two = 0
                }
            }

            fmt.Println("puntos player one", puntos_player_one)
            fmt.Println("puntos player two", puntos_player_two)
            fmt.Println("set ganados player one", set_ganados_player_one)
            fmt.Println("set ganados player two", set_ganados_player_two)

            if set_ganados_player_two >= 6 {
                fmt.Println("gano player two")
                break
            }

            if set_ganados_player_one >= 6 {
                fmt.Println("gano player one")
                break
            }

        }
    }
}

/*
   para mi un channel recibe la info y despues agregar una probabilidad de fallar
   un game y despues set

   a) uno arranca y se elige la probabilidad de que el saque sea bueno
   opciones de saque
       1 let
       2 out
       3 pasa
       4 ace

   sino tiene 2 out el punto es para el rival

   b) por cada vez que pasa (jugador recibe)
       1 no le pega
       2 le pega y no pasa
       3 le pega y pasa

   si (no le pega) o (no pasa) punto para el rival

   cuando llegan a 45 ganan el game cuando gana 6 y la difrencia de dos gana el set

*/
