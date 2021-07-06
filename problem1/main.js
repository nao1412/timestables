var START = 1;
var GOAL = 9;
let ESCAPE = [];
let nnn = [];
let mmm = [];

let times_tables = [];
let clm = [], idx = [];
for (let i = START; i <= GOAL; i++) {
  clm.push('x' + String(i))
  if (ESCAPE.includes(i)) continue;
  let TABLE = [];
  idx.push(String(i) + ' times')
  for (let j = START; j <= GOAL; j++) {
    if (nnn.length > 0) {
      for (let k = 0; k < nnn.length; k++) {
        if ((i * j) % nnn[k] == 0) {
          TABLE.push(String(i) + "x" + String(j) + "=" + String(i*j));
        } else {
          TABLE.push("");
        }
      }
    } else if (mmm.length > 0) {
      for (let k = 0; k < mmm.length; k++) {
        if ((i * j) % mmm[k] == 0) {
          TABLE.push("");
        } else {
          TABLE.push(String(i) + "x" + String(j) + "=" + String(i*j));
        }
      }
    } else {
      TABLE.push(String(i) + "x" + String(j) + "=" + String(i*j));
    }
  }
  times_tables.push(TABLE);
}

console.log(times_tables);
const transpose = a => a[0].map((_, c) => a.map(r => r[c]));
data = transpose(times_tables);
console.log(data);

// htmlで出力するときにi==0ならTIMESでj==0ならx1を出力。i==0andj==0なら空白