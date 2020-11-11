#include <stdio.h>
#include <time.h>

#define nbmove 8

#define SIDE 5
#define SQR_SIDE (SIDE * SIDE)

typedef unsigned int uint;

int movex[]={-1,-1,-2,-2,+1,+1,+2,+2};
int movey[]={-2,+2,-1,+1,+2,-2,+1,-1};
int shift[8];
int shift_0;
int shift_1;
int shift_2;
int shift_3;
int shift_4;
int shift_5;
int shift_6;
int shift_7;
int circuit[25];
int nbcoup=1;
int nbsol=0;

void solve(int nb, int x, int y);
void paint();

main() {
    clock_t tstart, tstop;
    double duration;
    int i=0,xs=0,ys=0;

    tstart=clock();

    printf("Cherche avec %d\n",SIDE);
    printf("ok");
    for (i=0; i<8; i++) {
        shift[i] = movex[i] * SIDE + movey[i];
    }
    shift_0=shift[0];
    shift_1=shift[1];
    shift_2=shift[2];
    shift_3=shift[3];
    shift_4=shift[4];
    shift_5=shift[5];
    shift_6=shift[6];
    shift_7=shift[7];

    nbsol=0;
    for (xs=0;xs<SIDE;xs++) {
        for (ys=0;ys<SIDE;ys++) {
            solve(1,xs,ys);
        }
    }
    tstop = clock();
    duration = (double)(tstop - tstart) / CLOCKS_PER_SEC;
    fprintf(stderr,"\n%dx%d cases nbsol=%d %.2f s\n",SIDE,SIDE,nbsol,duration);
    fflush(stdout);
    fflush(stderr);
}

void solve(int nb, int x, int y) {
    uint newx,newy;
    int pos;

    pos = x * SIDE + y;
    circuit[pos] = nb;
    if (nb == SQR_SIDE){
        //paint();
        nbsol += 1;
        circuit[pos]=0;
        return;
    }

    nb++;
    newx = x - 1;
    if (newx < SIDE) {
        newy = y - 2;
        if (newy < SIDE && circuit[pos + shift_0] == 0)
            solve(nb, newx, newy);
        newy = y + 2;
        if (newy < SIDE && circuit[pos + shift_1] == 0)
            solve(nb, newx, newy);
    }

    newx = x - 2;
    if (newx >= 0 && newx < SIDE) {
        newy = y - 1;
        if (newy < SIDE && circuit[pos + shift_2] == 0)
            solve(nb, newx, newy);
        newy = y + 1;
        if (newy < SIDE && circuit[pos + shift_3] == 0)
            solve(nb, newx, newy);
    }

    newx = x + 1;
    if (newx < SIDE) {
        newy = y + 2;
        if (newy < SIDE && circuit[pos + shift_4] == 0)
            solve(nb, newx, newy);
        newy = y - 2;
        if (newy < SIDE && circuit[pos + shift_5] == 0)
            solve(nb, newx, newy);
    }

    newx = x + 2;
    if (newx < SIDE) {
        newy = y + 1;
        if (newy < SIDE && circuit[pos + shift_6] == 0)
            solve(nb, newx, newy);
        newy = y - 1;
        if (newy < SIDE && circuit[pos + shift_7] == 0)
            solve(nb, newx, newy);
    }

    circuit[pos]=0;
}

void paint() {
    int x,y;
    printf("\n");
    for (x=0;x<SIDE;x++) {
        for (y=0;y<SIDE;y++) {
            if (SQR_SIDE<100) {
                printf("%02d ",circuit[x*SIDE+y]);
            } else {
                printf("%03d ",circuit[x*SIDE+y]);
            }
        }
        printf("\n");
    }
}
