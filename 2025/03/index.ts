import * as fs from 'fs';
import * as readline from 'readline';

let passwordOne = 0;
let passwordTwo = 0;
const filePath = 'data.txt';

async function run(): Promise<void> {
	try {
		const fileStream = fs.createReadStream(filePath);

		const rl = readline.createInterface({
			input: fileStream,
			crlfDelay: Infinity
		});

		rl.on('line', (line: string) => {
			partOne(line);
			partTwo(line);
		});

		await new Promise<void>((resolve) => {
			rl.on('close', () => {
				console.log(passwordOne);
				console.log(passwordTwo);
				resolve();
			});
		});
	} catch (error) {
		console.error("Error reading file:", error);
	}
}

function partOne(line: string): void {
	const values = line.split("").map(Number);

	let maxIndex = 0;

	for (let i = 0; i < values.length - 1; i++ ) {
		if (values[i]! > values[maxIndex]!) {
			maxIndex = i;
		}
	}

	let nextMaxIndex = maxIndex + 1;
	for (let i = nextMaxIndex; i < values.length; i++) {
		if (values[i]! > values[nextMaxIndex]!) {
			nextMaxIndex = i;
		}
	}

	passwordOne += Number(`${values[maxIndex]}${values[nextMaxIndex]}`);
}

function partTwo(line: string): void {
	const values = line.split("").map(Number);
	let maxIndex = 0;
	const digits: number[] = [];
	for (let digitsLeft = 11; digitsLeft >= 0; digitsLeft--) {
		for (let i = maxIndex; i < values.length - digitsLeft; i++) {
			if (values[i]! > values[maxIndex]!) {
				maxIndex = i;
			}
		}
		digits.push(values[maxIndex]!);
		maxIndex = maxIndex + 1;
	}

	passwordTwo += Number(digits.join(""));
}

run();

