let caught = 8*8; // let indicates this is a binding. followed by the name of the binding (caught)
// caught binding pointing to an int type -> referencing the output value

console.log(caught * caught); // 4096

// when binding a points to a value, it does not mean it is tied to it forever
let mood = 'light';
console.log(mood); // light
mood = 'dark'; // change binging value
console.log(mood); //dark

// think of bindings as tentacle rather than boxes, they don't contain values, they grasp them

// two bindings can refer to the same value
let debt = -150;
let myDebt = debt - 99;
console.log(`$${myDebt}`); // -249
myDebt += 55;
console.log(myDebt); // 194


// the words 'var' and 'const' can be used to create bindings similar to let
var name = "Felix";
const greeting = 'Hello ';
console.log(greeting + name);

// var -> short for variable, way bindings were declared in pre-2015 when let didnt exist, var behaive weirds, not used

// const stands for constant -> as in constant binding, which points to the same value for as long as it lives

// digis can be a part of binding names
let catch22 = 22;
console.log(catch22);

// Environment -> a collection of bindings and their values that exist at a given time is called an environment
// when the program starts up, the env is not empty.

// return values -> showing or writing text to the screen is called a side effect.
console.log(Math.max(2,7)); // gives back the greatest number
let num = 20;
let num2 = 55;
console.log(Math.max(num, num2) + 300); // 355

// Control flow -> programs run from top to bottom
// conditional execution is created with the if statements

function prompt(input){
    console.log(input);
}

let pickNum = Number(prompt("Pick a number"));

if(Number.isNaN(pickNum)){ // the isNaN is a function that returns true if it is not a number
    console.log('Sorry Not a Number!');
} else {
    console.log(pickNum + pickNum);
}

// you can turn that into a ternary
Number.isNaN(pickNum) ? console.log('Sorry Not a Number!') :  console.log(pickNum + pickNum); // produces the same output

// while loops
let ten = 1;
let counter = 0
while(counter < 10){
    ten *= 2;
    counter++;
}
console.log(ten); // 1024 -> power of 2

let result = 1;
for(let counter = 0; counter < 10; counter++){
    result *= 2;
}
console.log(result);

// breaking out of a loop
for(let current = 20; ; current++){ // notice we left the middle ; ; empty
    if(current % 7 == 0){
        console.log(current); // 21
        break; // break keyword to get out of the loop;
    }
}

// the '%' is great way to check if a number is divisible by another number. if it is the remainder of their division is zero

switch(prompt("Rainy")){
    case 'Rainy':
        console.log('It is raining');
        break;
    case 'Sunny':
        console.log('Sunny outside');
        break;
    default:
        console.log('Nothing');
        break;
}

let hash = '#';
for(let counter = '#'; counter.length < 10; counter += '#'){
    console.log(counter);
}

for(let i = 0; i < 100; i++){
    if(i % 3 === 0 && i % 5 === 0){
        console.log('FizzBuzz');
    } else if(i % 3 === 0){
        console.log('Fizz');
    } else if (i % 5 === 0){
        console.log('Buzz');
    } else {
        console.log(i);
    }
}

