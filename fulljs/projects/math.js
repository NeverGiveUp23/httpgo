// the math object

// the Math object is used as a container to group a bunch of related functionality. There is only one Math object and it is almost never useful as a value
//Math object provides a namespace so that all these functions and values do not have global bindings

function randomPointOnCircle(radius){
    let angle = Math.random() * 2 * Math.PI;

    return {
        x: radius * Math.cos(angle),
        y: radius * Math.sin(angle)
    };
}

console.log(randomPointOnCircle(2));

console.log(Math.random());

// if we want a whole number and not a fractional one when we call Math.random, we can use Math.floor, which rounds down to the nearest whole number
console.log(Math.floor(Math.random() * 10));// multiplying by 10 gives us a number greater than or equal to 0 and below 10;


