// the main idea behind OOP is to use objects, or rather types of objects, as the unit of program organization


// lets start with methods
function speak(line){
    console.log(`The ${this.type} rabbit says ${line}`); // binding called this in its body automatically points at the object on which it was called.
};

let whiteRabbit = {
    type: "white",
    speak
};
let hungryRabbit = {
    type: "Black",
    speak
};

whiteRabbit.speak("Oh my fur got caught");
hungryRabbit.speak("Im starving");

speak.call(whiteRabbit, "Now"); // using the call method takes 'this' value as is first argument and treats further arguments as normal params


// each function has its own 'this' binding
// arrow functions are different, they do not bind 'this' but can see the 'this' binding of the scope around them.
// thus, you can do something like the following code which references 'this' from inside a local function

let finder = {
    find(array){
        return array.some(v => v == this.value);
    },
    value: 5
};
// A property like find(array) in an object expression is a shorthand way of defining a method.
// It creates a property called find and gives it a function as its value.

console.log(finder.find([4 , 5]));


// PROTOTYPES
// it wouldbe nice to keep the types methods in a single place, rather tan adding them to each object individually
// in JS, prototypes do exactly that

// objects can be linked to other objects to do that.
let empty = {};
console.log(empty.toString);
// -> [Function: toString]
console.log(empty.toString());
// -> [object Object]

// it looks like we just pulled a property out of an empty object. In fact toString is a method stored in Object.prototype, meaning it is avail. in most objects
//When an object gets a request for a property that it doesn’t have, its prototype will be searched for the property.
console.log(Object.getPrototypeOf({}) == Object.prototype);
// -> true
console.log(Object.getPrototypeOf(Object.prototype));
// -> null
//As you’d guess, Object.getPrototypeOf returns the prototype of an object.

console.log(Object.getPrototypeOf(Math.max) == Function.prototype);
// -> true
console.log(Object.getPrototypeOf([]) == Array.prototype);
// -> true


// You can use Object.create to create an object with a specific prototype.
let protoRabbit = {
    speak(line) {
        console.log(`The ${this.type} rabbit say's ${line}`);
    },
    run(speed){
        console.log(`The ${this.type} runs ${speed}`);
    }
};
let blackRabbit = Object.create(protoRabbit);
blackRabbit.type = "Black";
blackRabbit.speak("Hello");
blackRabbit.run("fast");
// -> The Black rabbit say's Hello
// -> The Black runs fast
// the protoRabbit acts like a container for the properties shared by all rabbits.


// CLASSES
class Rabbit { // the class keyword starts the declaration
    constructor(type) {
        this.type = type;
    }
    speak(line){
        console.log(`The ${this.type} rabbit say's ${line}`);
    }
    run(speed){
        console.log(`The ${this.type} runs ${speed}`);
    }
}

let killerRabbit = new Rabbit("Killer");
killerRabbit.speak("i will kill you");


class Particle{
    speed = 0;
    constructor(position) {
        this.position = position;
    }
}


// private properties
class SecretObject {
    #getSecret(){
        return "I ate all the fruit";
    };

    interrogate(){
        let shallISayIt = this.#getSecret();
        return "never";
    };
}

// private members need to be declared in the class declaration to be available to all
class RandomSource {
    #max;
    constructor(max) {
        this.#max = max;
    };

    getNumberMax(){
        return Math.floor(Math.random() * this.#max);
    }
    getOfficialNumber(){
        return this.#max;
    }
}


let random = new RandomSource(22);
console.log(random.getNumberMax()); // only way to get the max number
console.log(random.getOfficialNumber()); // -> 22

// When you add a property to an object, whether it is present in the prototype or not, the property is added to the object itself.
// If there was already a property with the same name in the prototype, this property will no longer affect the object,
// now hidden behind the object’s own property.

Rabbit.prototype.teeth = "small";
console.log(killerRabbit.teeth);
// -> small
killerRabbit.teeth = "Long and sharp"; // takes the new effect of the teeth object
console.log(killerRabbit.teeth);
// -> Long and sharp
console.log(new Rabbit("basic").teeth);
// -> small

// overriding properties that exist in the prototype can be useful, where in the 'teeth' example we can express a different value, whereas
// a generic class of rabbit can sta with the default standard aka 'small' in this case

// getters, setters
class Temperature {
    constructor(celsius) {
        this.celsius = celsius;
    }

    get fahrenheit(){
        return this.celsius * 1.8 + 32;
    }

    set fahrenheit(value){
        this.celsius = (value - 32) / 1.8;
    }

    static fromFahrenheit(value){
        return new Temperature((value - 32) / 1.8);
    }
};

let temp = new Temperature(22);
console.log(temp.fahrenheit);
// -> 71.6

class Dog extends Rabbit { // Dog is a subclass to the superclass Rabbit
    #owner;
    constructor(type){
        super(type);
        this.#owner = super.owner;
    };

    get owner(){
       return this.#owner;
    }
}

let dog = new Dog("Dog");
dog.speak("woof");


