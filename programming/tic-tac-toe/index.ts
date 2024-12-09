const board = [
 "1",
 "2",
 "3",
 "4",
 "5",
 "6",
 "7",
 "8",
 "9",
]

const O = "O"
const X = "X"

let currentTurn = X

console.log("Welcome to Tic Tac Toe game!")

  process.stdout.write(`${board[6]}, ${board[7]}, ${board[8]}\n`);
  process.stdout.write(`${board[3]}, ${board[4]}, ${board[5]}\n`);
  process.stdout.write(`${board[0]}, ${board[1]}, ${board[2]}\n`);

console.log(`${currentTurn} turn to play, select a number from 1-9`)

const winConditions=[
  [0,1,2],
  [2,3,4],
  [5,6,7],
  [0,3,6],
  [1,4,7],
  [0,4,8],
  [6,4,2],
]

const checkWin = (board: string[]) => {
  for (const condition of winConditions ) {
    const [c0, c1, c2] = condition
    if (board[c0] === board[c1] && board[c1] === board[c2]) {
      return true
    }
  }

  return false
}

const checkDraw = (board: string[]) =>{
  for(const place of board){
    if(![O,X].includes(place)) return false
  }

  return true
}

for await (const line of console) {
  if ([X,O].includes(board[parseInt(line) - 1])) {
    console.log(`Position already ocupied, tried another one`)
    continue
  }

  board[parseInt(line) - 1] = currentTurn


  process.stdout.write(`${board[6]}, ${board[7]}, ${board[8]}\n`);
  process.stdout.write(`${board[3]}, ${board[4]}, ${board[5]}\n`);
  process.stdout.write(`${board[0]}, ${board[1]}, ${board[2]}\n`);

  if (checkWin( board )) {
    console.log(`${currentTurn} won!`)
    break
  }

  if (checkDraw( board )) {
    console.log(`Match drawn`)
    break
  }

  currentTurn =  currentTurn === X ? O : X
  console.log(`${currentTurn} turn to play, select a number from 1-9`)
}
