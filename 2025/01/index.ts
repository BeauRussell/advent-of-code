import * as fs from 'fs';
import * as readline from 'readline';

let pointer = 50;
let password = 0;
const filePath = 'data.txt';

async function run(): Promise<void> {
	try {
		const fileStream = fs.createReadStream(filePath);

		const rl = readline.createInterface({
			input: fileStream,
			crlfDelay: Infinity
		});

		rl.on('line', (line: string) => {
			processInstruction(line);
		});

		await new Promise<void>((resolve) => {
			rl.on('close', () => {
				console.log(password);
				resolve();
			});
		});
	} catch (error) {
		console.error("Error reading file:", error);
	}
}

function processInstruction(line: string): void {
	const direction = line[0]; 
	const distance = Number(line.substring(1));
	const originalPointer = pointer;
	const numToRun = Math.floor(distance / 100) + 1;

	if (numToRun > 1) console.log('Over 100');
	for (let i = 1; i <= numToRun; i++) {
		if (i < numToRun) {
			password++;
			continue;
		}
		const turns = distance % 100;
		pointer = direction === 'R' ? wrapTo99(pointer + turns) : wrapTo99(pointer - turns);
		if (direction === 'R' && pointer < originalPointer) password++;
		else if (direction === 'L' && pointer > originalPointer) password++; 
	}
	return;
}

function wrapTo99(currentNumber: number): number {
	const max = 100;
	let wrapped = currentNumber % max;

	if (wrapped < 0) {
		wrapped += max;
	}
	
	return wrapped;
}

run();
