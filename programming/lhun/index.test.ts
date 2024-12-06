import { expect, test } from "bun:test";
import { verifyLhun } from "./index.ts";

test("Returns true for valid digit", () => {
  expect(verifyLhun("17893729974")).toBe(true);
});

test("Returns false for invalid digit", () => {
  expect(verifyLhun("17893729978")).toBe(false);
});
