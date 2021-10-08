function main() {
    let input = document.getElementById("input").value

    let alphabet = []
    createAlphabet(65, alphabet)

    let result = checkLetters(0, input.toUpperCase(), alphabet)

    document.getElementById("result").innerHTML = result
}

function checkLetters(n, input, alphabet) {
    if (alphabet.length == 0) return "Ja, Pangram!"
    if (n >= input.length) return "Kein Pangram!"
    if (alphabet.includes(input[n])) alphabet.splice(alphabet.indexOf(input[n]), 1)

    return checkLetters(n + 1, input, alphabet)
}

function createAlphabet(n, alphabet) {
    if (n > 90) {
        return
    }
    alphabet.push(String.fromCharCode(n).toUpperCase())

    createAlphabet(n + 1, alphabet)
}