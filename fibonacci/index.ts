import express, { Request, Response } from "express";

// Initialize express app
const app = express();
const port = 8080;

// Middleware to parse JSON body
app.use(express.json());

/**
 * Generate Fibonacci sequence up to n elements
 * @param n Number of elements to generate
 * @returns Array containing the Fibonacci sequence
 */
const fibonacci = (n: number): number[] => {
  if (n <= 0) return [];
  if (n === 1) return [0];
  if (n === 2) return [0, 1];

  const result: number[] = [0, 1];
  for (let i = 2; i < n; i++) {
    result.push(result[i - 1] + result[i - 2]);
  }
  return result;
};

/**
 * Fibonacci sequence endpoint handler
 */
app.get("/", async (req: Request, res: Response) => {
  try {
    // Get the limit from query parameters or use default 10
    const limit = parseInt(req.query.limit as string) || 10;

    // Validate input
    if (isNaN(limit)) {
      return res.status(400).json({
        status: "error",
        message: "Invalid limit parameter. Please provide a number.",
      });
    }

    if (limit <= 0) {
      return res.status(400).json({
        status: "error",
        message: "Limit must be a positive number.",
      });
    }

    // Generate Fibonacci sequence
    const sequence = fibonacci(limit);

    // Return the result
    return res.status(200).json({
      status: "success",
      count: sequence.length,
      sequence: sequence,
    });
  } catch (error) {
    console.error("Error processing Fibonacci request:", error);
    return res.status(500).json({
      status: "error",
      message: "Internal server error",
    });
  }
});

// Start the server
app.listen(port, () => {
  console.log(`Fibonacci service listening on port ${port}`);
});
