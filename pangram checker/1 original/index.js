function main() {
    let input = document.getElementById("input").value
    let result = checkLetters(input.toUpperCase())
    document.getElementById("result").innerHTML = result
}

function checkLetters(input) {
    let alphabet = []
    for (let c = 65; c <= 90; c++) {
        alphabet.push(String.fromCharCode(c).toUpperCase())
    }
    for (let char of input) {
        if (alphabet.length == 0)
            return "Ja, Pangram!"
        if (alphabet.includes(char)) alphabet.splice(alphabet.indexOf(char), 1)
    }
    return "Kein Pangram!"
}
