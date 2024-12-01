import * as fs from 'fs';
import * as readline from 'readline';
import {ReadStream} from "node:fs";
import {Interface} from "node:readline";

function quickSort(arr: string[]): string[] {
    if (arr.length <= 1) {
        return arr;
    }

    const pivot: string = arr[arr.length - 1];
    const leftArr: string[] = [];
    const rightArr: string[] = [];

    for (let i: number = 0; i < arr.length - 1; i++) {
        if (arr[i] < pivot) {
            leftArr.push(arr[i]);
        } else {
            rightArr.push(arr[i]);
        }
    }

    return [...quickSort(leftArr), pivot, ...quickSort(rightArr)];
}

async function readLines(filePath: string): Promise<string[][]> {
    const fileStream: ReadStream = fs.createReadStream(filePath);
    const array: string[][] = [[],[]];

    const rl: Interface = readline.createInterface({
        input: fileStream
    });

    for await (const line of rl) {
        const lineArr: string[] = line.split('   ');
        array[0].push(lineArr[0]);
        array[1].push(lineArr[1]);
    }

    return array;
}

function checkTotalDistance(left: string[], right: string[]): number {
    let distance: number = 0;
    for (let i: number = 0; i < left.length; i++) {
        distance += Math.abs(parseInt(left[i]) - parseInt(right[i]));
    }

    return distance
}

function checkSimilarityScore(left: string[], right: string[]): number {
    let similarity: number = 0;
    let current: string = '0';
    for (let i: number = 0; i < right.length; i++) {
        let appearances: number = 0;
        current = left[i];
        let j: number = 0;
        while (true) {
            if (right[j] === current) {
                appearances += 1;
            } else if (right[j] > current) {
                break;
            }
            j = j + 1;
        }

        similarity += appearances * parseInt(current);
    }

    return similarity;
}

async function main(): Promise<void> {
    const arrays: string[][] = await readLines('../../puzzles/01/puzzle.txt');
    const left: string[] = quickSort(arrays[0]);
    const right: string[] = quickSort(arrays[1]);

    console.log(checkTotalDistance(left, right));
    console.log(checkSimilarityScore(left, right));
}

main();
