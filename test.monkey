let a = 1;
let b = 2;
let c = 3;
a;
b;
c;
let sum = fn(x, y) { return x + y; };
sum(a, b);
sum(b, c);
let total = sum(a, sum(b, c));
total;
let makeGreeter = fn(greeting) { fn(name) { greeting + " " + name + "!" } };
let hello = makeGreeter("Hello");
hello("manish")
let heythere = makeGreeter("Hey there");
heythere("manish")
