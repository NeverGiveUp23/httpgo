// when you arent sure whether a given value produces an object, but still want to read a property from it when it does, use the ?. notation

function isObject(object){
    return object.address?.city;
}

console.log(isObject({
    address: {city: "Toronto"}
})); // -> Toronto

console.log(isObject({
    name: "Vera"
}))  // -> undefined
// The expression a?.b means the same as a.b when a isn’t null or undefined.

// JSON

// propeerties grasp their values rather than contain it, object and arrays are stored in the computers memory as sequences of bits holding
// the addresses--the place in memory--of their contents.

// we can serialize the data to send over -> convert into a flat descriptor. JSON is the popular choice

let string = JSON.stringify({
    "squirrel": false,
    "events": ["Food", "Fight", "Sleep"]
});
console.log(string); // {"squirrel":false,"events":["Food","Fight","Sleep"]}

console.log(JSON.parse(string).events); // returns only the properties in the events object


function repeat(n, actions){
    for(let i = 0; i < n; i++){
        actions(i);
    }
}

let labels = [];
repeat(5, i => {   // this type of notation provides a body written as a function value.
    labels.push(`Unit ${i+1}`);
});
console.log(labels); // [ 'Unit 1', 'Unit 2', 'Unit 3', 'Unit 4', 'Unit 5' ]


function countBy(items, groupName){
    let counts = [];
    for(let item of items){
        let name = groupName(item);
        let known = counts.find(c => c.name == name);
        if(!known){
            counts.push({name, count: 1})
        } else {
            known.count++;
        }
    }
    return counts;
}

console.log(countBy([1,2,3,4,5, 6], n => n > 4));

// the counts function expects a collection(anything that we can loop over with)
