using System.Security.Cryptography.X509Certificates;
using System.Text.RegularExpressions;
using Xunit.Abstractions;

namespace AdventOfCode.Days;

public partial class DayThree
{
    private const string InputData = @"Data/DayThree";
    private readonly ITestOutputHelper _output;

    public DayThree(ITestOutputHelper output)
    {
        _output = output;
        _output.WriteLine("Day Three");
    }

    [Fact]
    public void PartOne()
    {
        var solution = 0;
        foreach (var line in File.ReadLines(InputData))
        {
            var matchGroups = MultiplicationRegex().Matches(line).Select(match => match.Groups[1]);
            foreach (var matchGroup in matchGroups)
            {
                solution += matchGroup.Value.Split(",").Select(int.Parse).Aggregate((x, y) => x * y);
            }
        }

        _output.WriteLine($"\tPart One: {solution}");
        Assert.Equal(170068701, solution);
    }

    [Fact]
    public void PartTwo()
    {
        var solution = 0;
        var conditionAlive = true;
        foreach (var line in File.ReadLines(InputData))
        {
            var matchGroups = ConditionalMultiplicationRegex().Matches(line).Cast<Match>().SelectMany(match => match.Groups.Cast<Group>().Skip(1));
            foreach (var matchGroup in matchGroups)
            {
                _output.WriteLine(matchGroup.Value);
                if (matchGroup.Value is "do()")
                {
                    conditionAlive = true;
                }
                else if (matchGroup.Value is "don't()")
                {
                    conditionAlive = false;
                }
                else if (conditionAlive && !string.IsNullOrWhiteSpace(matchGroup.Value))
                {
                    solution += matchGroup.Value.Split(",").Select(int.Parse).Aggregate((x, y) => x * y);
                }
            }
        }

        _output.WriteLine($"\tPart Two: {solution}");
        Assert.Equal(78683433, solution);
    }

    [GeneratedRegex(@"mul\((\d+,\d+)\)")]
    private static partial Regex MultiplicationRegex();

    [GeneratedRegex(@"mul\((\d+,\d+)\)|(do\(\))|(don't\(\))")]
    private static partial Regex ConditionalMultiplicationRegex();
}