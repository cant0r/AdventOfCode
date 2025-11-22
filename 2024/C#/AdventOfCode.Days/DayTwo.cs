using Xunit.Abstractions;

namespace AdventOfCode.Days;

public class DayTwo
{
    private const string InputData = @"Data/DayTwo";
    private readonly ITestOutputHelper _output;

    public DayTwo(ITestOutputHelper output)
    {
        _output = output;
        _output.WriteLine("Day Two");
    }

    [Fact]
    public void PartOne()
    {
        var safeReports = 0;

        foreach (var line in File.ReadLines(InputData))
        {
            var levels = line.Split().Select(int.Parse).ToList();
            var ascendingOrderedLevels = levels.Order();
            var descendingOrderedLevels = levels.OrderDescending();

            if (levels.SequenceEqual(ascendingOrderedLevels) || levels.SequenceEqual(descendingOrderedLevels))
            {
                if (IsReportSafe(levels))
                {
                    safeReports++;
                }
            }
        }

        _output.WriteLine($"\tPart One: {safeReports}");
        Assert.Equal(326, safeReports);
    }

    [Fact]
    public void PartTwo()
    {
        var safeReports = 0;

        foreach (var line in File.ReadLines(InputData))
        {
            var levels = line.Split().Select(int.Parse).ToList();

            for (var removeIndex = 0; removeIndex < levels.Count; removeIndex++)
            {
                var cutLevels = levels.ToList();
                cutLevels.RemoveAt(removeIndex);
                var ascendingOrderedLevels = cutLevels.Order();
                var descendingOrderedLevels = cutLevels.OrderDescending();

                if (cutLevels.SequenceEqual(ascendingOrderedLevels) || cutLevels.SequenceEqual(descendingOrderedLevels))
                {
                    if (IsReportSafe(cutLevels))
                    {
                        safeReports++;
                        break;
                    }
                }
            }
        }

        _output.WriteLine($"\tPart Two: {safeReports}");
        Assert.Equal(381, safeReports);
    }

    private bool IsReportSafe(List<int> levels)
    {
        for (var i = 1; i < levels.Count; i++)
        {
            var difference = Math.Abs(levels[i] - levels[i - 1]);

            if (difference is < 1 or > 3)
            {
                return false;
            }
        }
        return true;
    }
}