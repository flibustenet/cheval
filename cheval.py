from array import array
try:
    xrange
except NameError:
    xrange=range
import sys
SIDE = 5
SQR_SIDE = SIDE * SIDE

circuit = array("l", [0]) * SQR_SIDE
nsolutions = 0

movex = [-1,-1,-2,-2,+1,+1,+2,+2]
movey = [-2,+2,-1,+1,+2,-2,+1,-1]
shift = [x * SIDE + y for x,y in zip(movex, movey)]
print shift

def showCircuit():
    global nsolutions
    sys.stdout.write('%d\n' % nsolutions)
    for x in xrange(SIDE):
        x_SIDE = x * SIDE
        for y in xrange(SIDE):
            if SQR_SIDE < 100:
                sys.stdout.write("%02d " % circuit[x_SIDE + y])
            else:
                sys.stdout.write("%03d " % circuit[x_SIDE + y])
        sys.stdout.write('\n')

def solve(nb, x, y,
        SIDE=SIDE, SQR_SIDE=SQR_SIDE, circuit=circuit,
        shift_0=shift[0],
        shift_1=shift[1],
        shift_2=shift[2],
        shift_3=shift[3],
        shift_4=shift[4],
        shift_5=shift[5],
        shift_6=shift[6],
        shift_7=shift[7],
        ):
    global nsolutions

    pos = x * SIDE + y
    circuit[pos] = nb
    #showCircuit()
    if nb == SQR_SIDE:
    #    showCircuit()
        nsolutions += 1
        circuit[pos] = 0
        return

    nb += 1
    newx = x - 1
    if newx >= 0 and newx < SIDE:
        newy = y - 2
        if newy >= 0 and newy < SIDE and not circuit[pos + shift_0]:
            solve(nb, newx, newy)

        newy = y + 2
        if newy >= 0 and newy < SIDE and not circuit[pos + shift_1]:
            solve(nb, newx, newy)

    newx = x - 2
    if newx >= 0 and newx < SIDE:
        newy = y - 1
        if newy >= 0 and newy < SIDE and not circuit[pos + shift_2]:
            solve(nb, newx, newy)

        newy = y + 1
        if newy >= 0 and newy < SIDE and not circuit[pos + shift_3]:
            solve(nb, newx, newy)


    newx = x + 1
    if newx >= 0 and newx < SIDE:
        newy = y + 2
        if newy >= 0 and newy < SIDE and not circuit[pos + shift_4]:
            solve(nb, newx, newy)

        newy = y - 2
        if newy >= 0 and newy < SIDE and not circuit[pos + shift_5]:
            solve(nb, newx, newy)

    newx = x + 2
    if newx >= 0 and newx < SIDE:
        newy = y + 1
        if newy >= 0 and newy < SIDE and not circuit[pos + shift_6]:
            solve(nb, newx, newy)

        newy = y - 1
        if newy >= 0 and newy < SIDE and not circuit[pos + shift_7]:
            solve(nb, newx, newy)


    circuit[pos] = 0

def main():
    sys.stdout.write("Search for side=%d\n" % SIDE)
    for x in xrange(SIDE):
        for y in xrange(SIDE):
            solve(1, x, y);
    #solve(1,0,0)
    sys.stdout.write("\n%dx%d case, %d solutions.\n" % (SIDE, SIDE, nsolutions))

#import psyco; psyco.full()
import time
s=time.time()
main()
sys.stdout.write("%f s\n" % (time.time()-s))
