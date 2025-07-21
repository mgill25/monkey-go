let x = 5;
let y = 10;
let add = fn(a, b) { a + b };
let result = add(x, y);
result;

let factorial = fn(n) {
    if (n < 2) {
        1
    } else {
        n * factorial(n - 1)
    }
};

factorial(5);

let greeting = "Hello, Monkey!";
greeting;