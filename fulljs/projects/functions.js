const square = function(x){
    return (x * x);
};

console.log(square(5));

// function is created with the keyword function. functions can have a set of params -> in the above case it is 'x'
// a function body contains the statements that are to be executed when the function is called

const roundTo = function(n, step){
    let remainder = n % step;
    return n - remainder + (remainder < step / 2 ? 0 : step);
}

console.log(roundTo(67, 10));

// each binding has a scope, which in theprogram is visible. bindings defined outside the function, block, or module, the
// scope is the whole program. called global scope like we learned in c/C++

// inside the scope are local bindings or local scope

let x = 90; // global
function getX(){
    let x = 10; // local scope to the function;
    return x;
}

console.log(x); // global scope value
console.log(getX()); // local scope value

// bindings created with let and const are local to the block in which they are defined.
// bindings with those keywords made inside a loop , the code before and after the loop will not 'see' it

// var in 2015 JS allows it to be global even inside the scope
if (true){
    let y = 20; // local scope
    var z = 70; // global scope // not used likely
}


// nested scope -> creating block scopes inside each other
const hummus = function(factor){
    const ingredient = function(amount, unit, name){
        let ingredientAmount = amount * factor; // can see the factor param from inside the nested function
        if(ingredientAmount > 1){
            unit += "s";
        }
        console.log(`${ingredientAmount} ${unit} ${name}`);
        // outer scope function cant see the inner scope param from the inner function
    };
    ingredient(1, "can", "chickpeas");
    ingredient(2, "cup", "tahini");
    ingredient(3, "cup", "lemon juice");
    ingredient(4, "clove", "garlic");
    ingredient(5, "tablespoon", "olive oil");
};

hummus(5);
//5 cans chickpeas
// 10 cups tahini
// 15 cups lemon juice
// 20 cloves garlic
// 25 tablespoons olive oil

// things inside the ingredient function can see the factor binding from the outer function.
// but its local scope param, such as unit and ingredientAmount are not visible to the outer function.

// local scope can see all local scopes tha contain it, and all scopes can see global scopes
// this is called lexical scoping

// functions as values -> acts as a name for a piece of program


// functions defined below the code that uses it, function declarations are not part of the regular top-to-bottom flow of control.
console.log(missle());

function missle(){
    return "now";
}
// they are moved to the top of the scope and can be used within the program by all the code.


// ARROW FUNCTIONS
// the third notation for functions it uses (=>) arrow.
const roundToo = (n, step) => {
    let remainder = n % step;
    return n - remainder + (remainder < step / 2 ? 0 : step);
};

// the arrow function comes after the list of params and is followed by the functions body

// when there are only one param name, you can omit the parenthesis around the param list
// if the body is a single expression rather than a code block in braces, that expression will be returned from the function.

// these two functions do the samething
const squaretoo = (x) => {return x * x; };
const squeareA = (x) => x * x;

// when the arrow function has no params at all the param list is just empty
const emptyParam = () => {
    console.log('Empty')
};

const list = (list) => {
    const type = (name, age) => {
        let listType = age += 5;
        console.log(`${list}: ${name} ${listType}`);
    };
    type("Felix", 12);
}

list(1);
// arrow functions were added to mke smaller function look great



// THE CALL STACK
// because a function has to jump back to the place that called it when it returns, the computer must remember the context from which the
// call happened. the place where the computer stores this is called the 'call stack'

// everytime a function is called, the current context is stored on top of this stack. when returned, it removes the top context from the stack
// and uses that context to continue execution
// STORING THIS STACK REQUIRES SPACE IN THE COMPUTER MEMORY.
// when the stack grows to big you will get an error 'out of stack space'

// CLOSURE

function wrapValue(n) {
    let local = n;
    return () => local;
};
let wrapValue1 = wrapValue(5);
let wrapValue2 = wrapValue(4);
console.log(wrapValue1());
console.log(wrapValue2());
// functions accessing its local bindings

// shows that local bindings are created anew for every call and different calls do not affect each other's local bindings
// this feature is called 'closure'

function multiply(factor) {
    return number => number * factor;
};
let twice = multiply(4);
console.log(twice(3));

// multiply creates an environment in which its factor param is bound to 2.
// function value which is stored in twice remembers this env so that when it is called, it multiplies it by the argument.

// RECURSION

function power(base, exponent){
    if(exponent === 0){
        return 1;
    } else {
        return base * power(base, exponent - 1);
    };
};

console.log(power(5, 6));
// this implementation is a lot slower than a traditional for loop.





