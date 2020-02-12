#Revenge of the Pancakes


Run "go test" to see the demo

Formally: if we number the pancakes 1, 2, ..., N from top to bottom, you
choose the top i pancakes to flip. Then, after the flip, the stack is i, i­1, ..., 2, 1, i+1, i+2, ..., N. Pancakes 1,
2, ..., inow have the opposite side up, whereas pancakes i+1, i+2, ..., N have the same side up that they
had up before.

Input
The first line of the input gives the number of test cases, T. 
T test cases follow. Each consists of one line with a string S, each character of which is either 
+ (which represents a pancake that is initially happy side up) or 
­ (which represents a pancake that is initially blank side up). The string, when read left to right,
represents the stack when viewed from top to bottom.

Output
For each test case, output one line containing Case #x: y, where x is the test case number (starting
from 1) and y is the minimum number of times you will need to execute the maneuver to get all the
pancakes happy side up.
Limits
1 ≤ T ≤ 100.
Every character in S is either + or ­.
Small dataset
1 ≤ length of S ≤ 10.
Large dataset
1 ≤ length of S ≤ 100.
Sample
Input
5
­
­+
+­
+++
­­+­

Output

Case #1: 1

Case #2: 1

Case #3: 2

Case #4: 0

Case #5: 3

In Case #1, you only need to execute the maneuver once, flipping the first (and only) pancake.
In Case #2, you only need to execute the maneuver once, flipping only the first pancake.
In Case #3, you must execute the maneuver twice. One optimal solution is to flip only the first pancake,
changing the stack to ­­, and then flip both pancakes, changing the stack to ++. Notice that you cannot
just flip the bottom pancake individually to get a one­move solution; every time you execute the
maneuver, you must select a stack starting from the top.
In Case #4, all of the pancakes are already happy side up, so there is no need to do anything.
In Case #5, one valid solution is to first flip the entire stack of pancakes to get +­++, then flip the top
pancake to get ­­++, then finally flip the top two pancakes to get ++++.

