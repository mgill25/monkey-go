// Hello, this is a single-line comment. :)
let x = 5;
let y = 10;
x + y;

let add = fn(a, b) { return a + b; };
add(x, y);

let max = fn(a, b) {
	if (a > b) {
		return a;
	} else {
		return b;
	}
};
max(x, y);

let factorial = fn(n) {
	if (n < 2) {
		return 1;
	} else {
		return n * factorial(n - 1);
	}
};
factorial(5);

if (x < y) {
	y
} else {
	x
};