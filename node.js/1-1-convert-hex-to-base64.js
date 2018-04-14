var src = new Buffer('49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d', 'hex')

console.log(Buffer.from(src, 'hex'))
// <Buffer 49 27 6d 20 6b 69 6c 6c 69 6e 67 20 79 6f 75 72 20 62 72 61 69 6e 20 6c 69 6b 65 20 61 20 70 6f 69 73 6f 6e 6f 75 73 20 6d 75 73 68 72 6f 6f 6d>

console.log(src.toString())
// I'm killing your brain like a poisonous mushroom

console.log(Buffer.from(src, 'hex').toString('base64'))
// SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
