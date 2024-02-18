package main

import (
    "fmt"
    "time"
    "math/rand"
)

const (
    width  = 20  // width of the field
    height = 10  // height of the field
)

// Define the structure for the ball
type ball struct {
    x, y int // ball position
    dx, dy int // ball direction
}

// Define the structure for the paddle
type paddle struct {
    y int // paddle position
}

func main() {
    // Initialize the ball at the center of the field with random movement
    var ball = ball{x: width / 2, y: height / 2, dx: 1, dy: 1}
    // Initialize paddles at the center of the field
    var left = paddle{height / 2}
    var right = paddle{height / 2}

    // Main game loop
    for {
        // Clear the terminal screen
        fmt.Print("\033[H\033[2J")

        // Update ball position
        ball.x += ball.dx
        ball.y += ball.dy

        // Check for collisions with walls
        if ball.y == 0 || ball.y == height - 1 {
            ball.dy = -ball.dy // reverse vertical direction
        }

        // Check for collisions with paddles
        if ball.x == 1 && ball.y == left.y {
            ball.dx = -ball.dx // reverse horizontal direction
        }
        if ball.x == width - 2 && ball.y == right.y {
            ball.dx = -ball.dx // reverse horizontal direction
        }

        // Check for scoring
        if ball.x == 0 || ball.x == width - 1 {
            // Reset ball to center and change its direction randomly
            ball = ball{x: width / 2, y: height / 2, dx: -ball.dx, dy: rand.Intn(3) - 1}
        }

        // Draw paddles
        for i := 0; i < height; i++ {
            if i == left.y || i == right.y {
                fmt.Print("X") // draw a paddle
            } else {
                fmt.Print(" ") // draw empty space
            }
        }
        fmt.Println()

        // Draw the ball
        for i := 0; i < height; i++ {
            for j := 0; j < width; j++ {
                if i == ball.y && j == ball.x {
                    fmt.Print("O") // draw the ball
                } else {
                    fmt.Print(" ") // draw empty space
                }
            }
            fmt.Println()
        }

        // Wait before the next frame
        time.Sleep(100 * time.Millisecond)
    }
}
