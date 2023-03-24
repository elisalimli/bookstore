import { get } from "..";

export async function getHello() {
  try {
    const response = await get<string>(`/`);

    return response.data;
  } catch (error) {
    console.error("getUsers - Error: ", error);
    throw error;
  }
}
