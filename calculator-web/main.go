package main

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"
)

func calculate(expr string) (int, error) {
    normalized := ""
    for _, r := range expr {
        if r == '+' || r == '-' || r == '*' || r == '/' {
            normalized += " " + string(r) + " "
        } else {
            normalized += string(r)
        }
    }

    parts := strings.Fields(normalized)
    if len(parts) != 3 {
        return 0, fmt.Errorf("invalid input")
    }

    left, err := strconv.Atoi(parts[0])
    if err != nil {
        return 0, err
    }
    right, err := strconv.Atoi(parts[2])
    if err != nil {
        return 0, err
    }

    switch parts[1] {
    case "+":
        return left + right, nil
    case "-":
        return left - right, nil
    case "*":
        return left * right, nil
    case "/":
        if right == 0 {
            return 0, fmt.Errorf("division by zero")
        }
        return left / right, nil
    default:
        return 0, fmt.Errorf("invalid operator")
    }
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
    expr := r.URL.Query().Get("expr")
    if expr == "" {
        http.Error(w, "Usage: /calc?expr=3+4", http.StatusBadRequest)
        return
    }

    result, err := calculate(expr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "Result: %d\n", result)
}

func main() {
    // Serve API
    http.HandleFunc("/calc", calcHandler)

    // Serve static UI
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    fmt.Println("Calculator web app running on port 8080...")
    http.ListenAndServe(":8080", nil)
}

