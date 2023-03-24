import { useQuery } from "react-query";
import { getHello } from "../hello/hello";

export function useHello() {
  return useQuery("hello", () => getHello());
}
