export const verifyLhun = (x: string): boolean => {
  const [payload, digit] = [x.slice(0, -1), x.slice(-1)];
  const s = payload
    .split("")
    .map((x) => parseInt(x))
    .reduce((acc, value, i) => {
      const sum = i % 2 === 1 ? value * 2 : value;
      if (sum > 9) return acc + sum - 9;
      return acc + sum;
    }, 0);

  return (10 - (s % 10)) % 10 === parseInt(digit);
};
