import * as fs from 'fs';
import * as readline from 'readline';

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
			processInput(line);
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

function processInput(line: string): void {
	const ranges = line.split(',');

	for (const range of ranges) {
		const arr = range.split('-');
		const min = Number(arr[0]);
		const max = Number(arr[1]);

		for (let i = min; i <= max; i++) {
			const string = i.toString();

			if (/^(\d+)\1+$/.test(string)) {
				password += i;
			}
		}
	}
}

run();
