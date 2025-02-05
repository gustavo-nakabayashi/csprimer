#define _GNU_SOURCE

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <time.h>
#include <sys/times.h>
#include <sched.h>

#define SLEEP_SEC 3
#define NUM_MULS 100000000
#define NUM_MALLOCS 100000
#define MALLOC_SIZE 1000



// TODO define this struct
struct profile_times {
  clock_t start_real;
  clock_t start_system;
  clock_t start_user;
  clock_t end_real;
  clock_t end_system;
  clock_t end_user;
};

// TODO populate the given struct with starting information
void profile_start(struct profile_times *t) {
  struct tms buf;

  t->start_real = times(&buf);
  t->start_system = buf.tms_stime;
  t->start_user = buf.tms_utime;
}

// TODO given starting information, compute and log differences to now
void profile_log(struct profile_times *t) {
  struct tms buf;

  t->end_real = times(&buf);
  t->end_system = buf.tms_stime;
  t->end_user = buf.tms_utime;

  double total_system, total_user, total_real;
  total_real = (double)(t->end_real - t->start_real) / sysconf(_SC_CLK_TCK);
  total_system = (double)(t->end_system - t->start_system) / sysconf(_SC_CLK_TCK);
  total_user = (double)(t->end_user - t->start_user) / sysconf(_SC_CLK_TCK);


  int cpu;

  cpu = sched_getcpu();

  printf("\033[0;31m");
  printf("[pid %d, cpu %d]", getpid(), cpu);

  printf("Real time: %f, User Time: %f, System time: %f\n", total_real, total_user , total_system);
}

int main(int argc, char *argv[]) {
  struct profile_times t;

  // TODO profile doing a bunch of floating point muls
  float x = 1.0;
  profile_start(&t);
  for (int i = 0; i < NUM_MULS; i++)
    x *= 1.1;
  profile_log(&t);

  // TODO profile doing a bunch of mallocs
  profile_start(&t);
  void *p;
  for (int i = 0; i < NUM_MALLOCS; i++)
    p = malloc(MALLOC_SIZE);
  profile_log(&t);

  // TODO profile sleeping
  profile_start(&t);
  sleep(SLEEP_SEC);
  profile_log(&t);
}


