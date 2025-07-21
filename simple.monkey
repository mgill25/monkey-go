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

let map = fn(arr, f) {
    let iter = fn(arr, acc) {
        if (len(arr) == 0) {
            acc
        } else {
            iter(rest(arr), push(acc, f(first(arr))));
        }
    };
    iter(arr, []);
};

let double = fn(x) { x * 2 };
map([1, 2, 3], double);

let reduce = fn(arr, initial, f) {
    let iter = fn(arr, result) {
        if (len(arr) == 0) {
            result
        } else {
            iter(rest(arr), f(result, first(arr)));
        }
    };
    iter(arr, initial);
};

let sum = fn(a, b) { a + b };
reduce([1, 2, 3, 4, 5], 0, sum);
