/* 
  (
    Root(),
    Projection(('name',)),
    Selection(('==', 'id', 5000)),
    Scan('movies'),
)

Root -> Projection -> Selection -> Scan

The pointer is the next

*/

import fs from "fs";

const table = [
  {
    name: "Foo",
    age: 28,
  },
  {
    name: "Bar",
    age: 48,
  },
  {
    name: "Baz",
    age: 8,
  },
];

class Scan {
  index: number;

  constructor() {
    this.index = 0;
  }

  self() {
    if (this.index >= table.length) {
      return null;
    }

    return table[this.index];
  }
}

// This will return data from db
class Root {
  result: Array<Object>;
  nextFunction: () => void;

  constructor(nextFunction: () => void) {
    this.result = [];
    this.nextFunction = nextFunction;
  }

  self() {
    return null;
  }

  next() {
    const n = this.nextFunction();

    if (n === null) {
      return this.result;
    }
  }
}

// This will be a function that returns only the keys filtered
// const Projection = (cols) => {};

const Instructions = [{


}];

const Executor = (instructions: Instructions) => {

};
