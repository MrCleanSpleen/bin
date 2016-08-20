import math as m

def find_collatz(x):
    seq = [x]
    while seq[-1] > 1:
       if x % 2 == 0:
         seq.append(x/2)
       else:
         seq.append(3*x+1)
       x = seq[-1]
    return seq

if __name__ == "__main__":
    find_collatz(7)
