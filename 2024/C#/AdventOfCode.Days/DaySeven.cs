using System.Globalization;
using Xunit.Abstractions;

namespace AdventOfCode.Days;

public class DaySeven
{
    private readonly ITestOutputHelper _output;
    private const string InputData = @"Data/DaySeven";

    public DaySeven(ITestOutputHelper output)
    {
        _output = output;
        _output.WriteLine("Day Seven");
    }

    [Fact]
    public void PartOne()
    {
        var solution = 0L;

        foreach (var expression in File.ReadLines(InputData))
        {
            var expressionParts = expression.Split(": ");
            var targetValue = long.Parse(expressionParts[0], CultureInfo.InvariantCulture);
            var values = expressionParts[1].Split().Select(long.Parse).ToArray();

            var isExpressionSolvable = IsExpressionSolvable(values, targetValue);

            if (isExpressionSolvable)
            {
                solution += targetValue;
            }
        }

        _output.WriteLine($"\tPart One: {solution}");
        Assert.Equal(303876485655L, solution);
    }

    [Fact]
    public void PartTwo()
    {
        var solution = 0L;

        foreach (var expression in File.ReadLines(InputData))
        {
            var expressionParts = expression.Split(": ");
            var targetValue = long.Parse(expressionParts[0], CultureInfo.InvariantCulture);
            var values = expressionParts[1].Split().Select(long.Parse).ToArray();

            var isExpressionSolvable = IsExpressionSolvable2(values, targetValue);

            if (isExpressionSolvable)
            {
                solution += targetValue;
            }
        }

        _output.WriteLine($"\tPart Two: {solution}");
        Assert.Equal(146111650210682L, solution);
    }

    private bool IsExpressionSolvable(long[] values, long targetValue) => IsExpressionSolvable(values[1..], targetValue, values[0]);

    private bool IsExpressionSolvable(long[] values, long targetValue, long accumulator)
    {
        if (values.Length == 0)
        {
            return accumulator == targetValue;
        }

        return IsExpressionSolvable(values[1..], targetValue, accumulator + values[0]) ||
               IsExpressionSolvable(values[1..], targetValue, accumulator * values[0]);
    }

    private bool IsExpressionSolvable2(long[] values, long targetValue) => IsExpressionSolvable2(values[1..], targetValue, values[0]);

    private bool IsExpressionSolvable2(long[] values, long targetValue, long accumulator)
    {
        if (values.Length == 0)
        {
            return accumulator == targetValue;
        }

        var concatenatedValues = long.Parse($"{accumulator}{values[0]}", CultureInfo.InvariantCulture);

        return IsExpressionSolvable2(values[1..], targetValue, accumulator + values[0]) ||
               IsExpressionSolvable2(values[1..], targetValue, accumulator * values[0]) ||
               IsExpressionSolvable2(values[1..], targetValue, concatenatedValues);
    }
}