import React from "react";
import { render, screen } from "@testing-library/react";
import App from "./App";

test("Given render App component When get by text learn react Then expect text to be in document", () => {
  // Given
  render(<App />);
  // When
  const linkElement = screen.getByText(/learn react/i);
  // Then
  expect(linkElement).toBeInTheDocument();
});
