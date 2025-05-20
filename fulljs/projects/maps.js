// Map(noun) is data structure that associates values (the keys) with others values. for example you might want to map names to ages
let ages = {
  Boris: 77,
  Mike: 99,
  Jess: 34,
};

console.log(`${ages.Boris}`);
// -> 77

// Object property names must be strings.
// If you need a map whose keys can’t easily be converted to strings—such as objects—you cannot use an object as your map.

// JavaScript comes with a class called Map
let names = new Map();
names.set("Boris", 77);
