using System.Globalization;
using Xunit.Abstractions;

namespace AdventOfCode.Days;

public class DayOne
{
    private const string InputData = @"Data/DayOne";
    private readonly ITestOutputHelper _output;

    public DayOne(ITestOutputHelper output)
    {
        _output = output;
        _output.WriteLine("Day One");
    }

    [Fact]
    public void Part_One()
    {
        var leftTeamIds = new List<int>();
        var rightTeamIds = new List<int>();

        foreach (var line in File.ReadLines(InputData))
        {
            var ids = line.Split("   ");
            leftTeamIds.Add(int.Parse(ids[0], CultureInfo.InvariantCulture));
            rightTeamIds.Add(int.Parse(ids[1], CultureInfo.InvariantCulture));
        }

        leftTeamIds.Sort();
        rightTeamIds.Sort();

        var solution = leftTeamIds.Zip(rightTeamIds, (l, r) => Math.Abs(l - r)).Sum();

        _output.WriteLine($"\tPart One: {solution}");
        Assert.Equal(2375403, solution);
    }

    [Fact]
    public void Part_Two()
    {
        var leftTeamIds = new List<int>();
        var rightTeamIds = new List<int>();

        foreach (var line in File.ReadLines(InputData))
        {
            var ids = line.Split("   ");
            leftTeamIds.Add(int.Parse(ids[0], CultureInfo.InvariantCulture));
            rightTeamIds.Add(int.Parse(ids[1], CultureInfo.InvariantCulture));
        }

        var solution = leftTeamIds.Sum(leftId => leftId * rightTeamIds.Count(rightId => rightId == leftId));

        _output.WriteLine($"\tPart Two: {solution}");
        Assert.Equal(23082277, solution);
    }
}