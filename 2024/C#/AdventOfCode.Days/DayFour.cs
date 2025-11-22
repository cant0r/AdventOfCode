using System.Text;
using System.Text.RegularExpressions;
using Xunit.Abstractions;

namespace AdventOfCode.Days;

public class DayFour
{
    private readonly ITestOutputHelper _output;
    private const string InputData = @"Data/DayFour";

    public DayFour(ITestOutputHelper output)
    {
        _output = output;
        _output.WriteLine("Day Four");
    }

    [Fact]
    public void PartOne()
    {
        var solution = 0;
        var wordMatrix = File.ReadLines(InputData).ToList();

        for (var i = 0; i < wordMatrix.Count; i++)
        {
            for (var j = 0; j < wordMatrix[i].Length; j++)
            {
                var startLetter = wordMatrix[i][j];

                if (startLetter is 'X')
                {
                    solution += CountMatches(wordMatrix, i, j, "XMAS");
                }
            }
        }

        _output.WriteLine($"\tPart One: {solution}");
        Assert.Equal(2401, solution);
    }

    [Fact]
    public void PartTwo()
    {
        var solution = 0;
        var wordMatrix = File.ReadLines(InputData).ToList();

        for (var i = 0; i < wordMatrix.Count; i++)
        {
            for (var j = 0; j < wordMatrix[i].Length; j++)
            {
                if (wordMatrix[i][j] is 'A')
                {
                    solution += CountMatchesTwo(wordMatrix, i, j);
                }
            }
        }

        _output.WriteLine($"\tPart Two: {solution}");
        Assert.Equal(1822, solution);
    }

    private int CountMatchesTwo(List<string> wordMatrix, int i, int j)
    {
        try
        {
            var pattern =
                string.Empty +
                wordMatrix[i - 1][j - 1] + wordMatrix[i - 1][j + 1] +
                wordMatrix[i + 1][j - 1] + wordMatrix[i + 1][j + 1];

            if (pattern is "MMSS" or "SSMM" or "SMSM" or "MSMS")
            {
                return 1;
            }
        }
        catch
        {
            return 0;
        }

        return 0;
    }

    private int CountMatches(List<string> wordMatrix, int row, int col, string pattern)
    {
        var matchBuilder = new StringBuilder();
        var wordMatrixWidth = wordMatrix[0].Length;
        var wordMatrixHeight = wordMatrix.Count;

        // Right
        try
        {
            matchBuilder.Append(
                string.Empty +
                wordMatrix[row][col] +
                wordMatrix[row][col + 1] +
                wordMatrix[row][col + 2] +
                wordMatrix[row][col + 3]);
        }
        catch
        {
        }

        // Left
        try
        {
            matchBuilder.Append(
                string.Empty +
                wordMatrix[row][col] +
                wordMatrix[row][col - 1] +
                wordMatrix[row][col - 2] +
                wordMatrix[row][col - 3]);
        }
        catch
        {
        }

        // Top
        try
        {
            matchBuilder.Append(
                string.Empty +
                wordMatrix[row][col] +
                wordMatrix[row - 1][col] +
                wordMatrix[row - 2][col] +
                wordMatrix[row - 3][col]);
        }
        catch
        {
        }

        // Bottom
        try
        {
            matchBuilder.Append(
                string.Empty +
                wordMatrix[row][col] +
                wordMatrix[row + 1][col] +
                wordMatrix[row + 2][col] +
                wordMatrix[row + 3][col]);
        }
        catch
        {
        }

        // Top Right
        try
        {
            matchBuilder.Append(
                string.Empty +
                wordMatrix[row][col] +
                wordMatrix[row - 1][col + 1] +
                wordMatrix[row - 2][col + 2] +
                wordMatrix[row - 3][col + 3]);
        }
        catch
        {
        }

        // Top Left
        try
        {
            matchBuilder.Append(
                string.Empty +
                wordMatrix[row][col] +
                wordMatrix[row - 1][col - 1] +
                wordMatrix[row - 2][col - 2] +
                wordMatrix[row - 3][col - 3]);
        }
        catch
        {
        }

        // Bottom Left
        try
        {
            matchBuilder.Append(
                string.Empty +
                wordMatrix[row][col] +
                wordMatrix[row + 1][col - 1] +
                wordMatrix[row + 2][col - 2] +
                wordMatrix[row + 3][col - 3]);
        }
        catch
        {
        }

        // Bottom Right
        try
        {
            matchBuilder.Append(
                string.Empty +
                wordMatrix[row][col] +
                wordMatrix[row + 1][col + 1] +
                wordMatrix[row + 2][col + 2] +
                wordMatrix[row + 3][col + 3]);
        }
        catch
        {
        }

        return Regex.Count(matchBuilder.ToString(), pattern);
    }
}