import urllib.request



def memoize(f):
    cache = {}

    def memoized(*arg, **vargs):
        if (arg in cache):
            return cache[arg]

        res = f(*arg, **vargs)
        cache[arg] = res
        return res

    return memoized

def fetch(url):
    with urllib.request.urlopen(url) as response:
        content = response.read().decode('latin-1')
        return content


@memoize
def fib(n):
    if n <= 1:
        return n
    return fib(n - 1) + fib(n - 2)


if __name__ == '__main__':
    print(fib(35))
    print(fib(35))
    print(fib(35))
