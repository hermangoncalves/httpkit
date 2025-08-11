import { createContext } from "react";

const initialState: ThemeProviderState = {
  theme: "system",
  setTheme: () => null,
};

export type Theme = "dark" | "light" | "system";


export type ThemeProviderState = {
  theme: Theme;
  setTheme: (theme: Theme) => void;
};
export const ThemeProviderContext =
  createContext<ThemeProviderState>(initialState);
