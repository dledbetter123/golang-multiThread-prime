# Prime Finder in Binary Files

This Go program efficiently searches for prime numbers in binary files using a multithreaded approach. It divides the file into segments, assigns them to worker threads, and processes these segments concurrently to maximize resource utilization and minimize processing time.

## Prerequisites

- A binary file to process (see the section on generating test data)

## Usage

1. **Compiling the Program**

   First, compile the program using the Go toolchain:

   ```bash
   go build -o primefinder main.go
   ```

   This command generates an executable named `primefinder`.

2. **Generating Test Data**

   The program includes a feature to generate a binary file filled with sequential numbers for testing purposes. To generate a file, run:

   ```bash
   ./primefinder -generate=true -min=<minimum value> -max=<maximum value> -rng=<number of values> -random=<random or sequential>
   ```

   By default, this generates a file named `newgen.dat`

3. **Running the Program**

   To run the program, use the following syntax:

   ```bash
   ./primefinder -pathname=<path to binary file> -M=<number of workers> -N=<segment size> -C=<chunk size>
   ```

   - `-pathname`: Path to the binary file to be processed.
   - `-M`: Number of worker threads to use for processing.
   - `-N`: Size of each segment in bytes. Each worker thread processes one segment at a time.
   - `-C`: Chunk size in bytes. Each read operation within a worker processes this amount of bytes.

   **Example:**

   ```bash
   ./primefinder -pathname=./data/newgen.dat -M=10 -N=65536 -C=1024
   ```

## Configuration Flags

### Number Generation

- `-generate`: Set to `true` to activate data file generation mode.
- `-min`:the minimum value (inclusive) that can be written to the file. Default is `0`.
- `-max`: the maximum value (inclusive) that can be written to the file. Default is `1,000,000`.
- `-rng`: how many values to write to the file.
- `-random`: whether the values should be generated randomly (`true`) or sequentially (`false`). Default is `true`.

### Usage

To generate a binary file with random numbers within a specified range, use the following command format:

```bash
go run main.go -generate -min=100 -max=1000000 -rng=500000 -random=false
Output:
Data file generated, min=100, max=1000000, samples=500000 Note the values generated are WITH replacement if random is enabled.
This was tested on ranges 1,000-1 billion to baseline accurate prime counting
```

### Prime Calculation

- `-pathname`: Required. Specifies the path to the binary file to process.
- `-M`: Number of worker threads (`1` by default).
- `-N`: Size of each segment in bytes (`65536` by default).
- `-C`: Chunk size in bytes for each read operation (`1024` by default).

## Output

The program logs the progress and results of the processing to the console. This includes the number of prime numbers found in each segment and the total number of primes found in the entire file. this is running on a (non-random) range from 0-500,000

```
(base) davidledbetter@Davids-MacBook-Pro-4 golang-multiThread-prime % go run main.go -pathname newgen.dat -M 50
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:0 Length:65536} with result: {Job:{Pathname:newgen.dat Start:0 Length:65536} Count:1028}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:65536 Length:65536} with result: {Job:{Pathname:newgen.dat Start:65536 Length:65536} Count:872}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:131072 Length:65536} with result: {Job:{Pathname:newgen.dat Start:131072 Length:65536} Count:825}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:262144 Length:65536} with result: {Job:{Pathname:newgen.dat Start:262144 Length:65536} Count:776}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:196608 Length:65536} with result: {Job:{Pathname:newgen.dat Start:196608 Length:65536} Count:787}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:327680 Length:65536} with result: {Job:{Pathname:newgen.dat Start:327680 Length:65536} Count:763}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:393216 Length:65536} with result: {Job:{Pathname:newgen.dat Start:393216 Length:65536} Count:763}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:458752 Length:65536} with result: {Job:{Pathname:newgen.dat Start:458752 Length:65536} Count:728}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:524288 Length:65536} with result: {Job:{Pathname:newgen.dat Start:524288 Length:65536} Count:739}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:589824 Length:65536} with result: {Job:{Pathname:newgen.dat Start:589824 Length:65536} Count:728}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:655360 Length:65536} with result: {Job:{Pathname:newgen.dat Start:655360 Length:65536} Count:718}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:720896 Length:65536} with result: {Job:{Pathname:newgen.dat Start:720896 Length:65536} Count:712}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:786432 Length:65536} with result: {Job:{Pathname:newgen.dat Start:786432 Length:65536} Count:712}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:851968 Length:65536} with result: {Job:{Pathname:newgen.dat Start:851968 Length:65536} Count:695}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:917504 Length:65536} with result: {Job:{Pathname:newgen.dat Start:917504 Length:65536} Count:708}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:983040 Length:65536} with result: {Job:{Pathname:newgen.dat Start:983040 Length:65536} Count:697}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:1048576 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1048576 Length:65536} Count:690}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:1114112 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1114112 Length:65536} Count:690}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:1179648 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1179648 Length:65536} Count:696}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:1245184 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1245184 Length:65536} Count:672}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:1310720 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1310720 Length:65536} Count:671}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:1376256 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1376256 Length:65536} Count:687}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:1441792 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1441792 Length:65536} Count:675}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:1507328 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1507328 Length:65536} Count:672}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:1572864 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1572864 Length:65536} Count:663}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:1638400 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1638400 Length:65536} Count:678}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:1703936 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1703936 Length:65536} Count:664}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:1769472 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1769472 Length:65536} Count:670}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:1835008 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1835008 Length:65536} Count:643}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:1900544 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1900544 Length:65536} Count:673}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:1966080 Length:65536} with result: {Job:{Pathname:newgen.dat Start:1966080 Length:65536} Count:663}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:2031616 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2031616 Length:65536} Count:642}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:2097152 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2097152 Length:65536} Count:672}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:2162688 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2162688 Length:65536} Count:646}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:2228224 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2228224 Length:65536} Count:655}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:2293760 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2293760 Length:65536} Count:634}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:2359296 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2359296 Length:65536} Count:641}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:2424832 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2424832 Length:65536} Count:651}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:2490368 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2490368 Length:65536} Count:667}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:2555904 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2555904 Length:65536} Count:657}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:2621440 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2621440 Length:65536} Count:642}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:2686976 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2686976 Length:65536} Count:639}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:2752512 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2752512 Length:65536} Count:648}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:2818048 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2818048 Length:65536} Count:637}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:2883584 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2883584 Length:65536} Count:642}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:2949120 Length:65536} with result: {Job:{Pathname:newgen.dat Start:2949120 Length:65536} Count:624}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:3014656 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3014656 Length:65536} Count:628}
2024/03/03 21:50:41 Worker 20 completed job: {Pathname:newgen.dat Start:3080192 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3080192 Length:65536} Count:652}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:3145728 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3145728 Length:65536} Count:626}
2024/03/03 21:50:41 Worker 2 completed job: {Pathname:newgen.dat Start:3211264 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3211264 Length:65536} Count:618}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:3276800 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3276800 Length:65536} Count:634}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:3342336 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3342336 Length:65536} Count:638}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:3407872 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3407872 Length:65536} Count:642}
2024/03/03 21:50:41 Worker 20 completed job: {Pathname:newgen.dat Start:3473408 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3473408 Length:65536} Count:633}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:3538944 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3538944 Length:65536} Count:631}
2024/03/03 21:50:41 Worker 2 completed job: {Pathname:newgen.dat Start:3604480 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3604480 Length:65536} Count:611}
2024/03/03 21:50:41 Worker 30 completed job: {Pathname:newgen.dat Start:3670016 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3670016 Length:65536} Count:612}
2024/03/03 21:50:41 Worker 2 completed job: {Pathname:newgen.dat Start:3997696 Length:2304} with result: {Job:{Pathname:newgen.dat Start:3997696 Length:2304} Count:19}
2024/03/03 21:50:41 Worker 45 completed job: {Pathname:newgen.dat Start:3735552 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3735552 Length:65536} Count:647}
2024/03/03 21:50:41 Worker 20 completed job: {Pathname:newgen.dat Start:3866624 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3866624 Length:65536} Count:619}
2024/03/03 21:50:41 Worker 11 completed job: {Pathname:newgen.dat Start:3801088 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3801088 Length:65536} Count:637}
2024/03/03 21:50:41 Worker 43 completed job: {Pathname:newgen.dat Start:3932160 Length:65536} with result: {Job:{Pathname:newgen.dat Start:3932160 Length:65536} Count:636}
Total primes: 41538

```