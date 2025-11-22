using Xunit.Abstractions;

namespace AdventOfCode.Days;

internal enum Direction
{
    Up = 0,
    Right,
    Down,
    Left,
}

internal class Guard((int Row, int Col) position, Direction direction, int progress)
{
    public (int Row, int Col) Position { get; set; } = position;
    public Direction Direction { get; set; } = direction;
    public int Progress { get; set; } = progress;
}

public class DaySix
{
    private readonly ITestOutputHelper _output;
    private const string InputData = @"Data/DaySix";

    public DaySix(ITestOutputHelper output)
    {
        _output = output;
        _output.WriteLine("Day Six");
    }

    [Fact]
    public void PartOne()
    {
        var solution = 0;
        var map = File.ReadAllLines(InputData).ToList();
        var seenPositions = new HashSet<(int Row, int Col)>();
        var guard = new Guard((0, 0), Direction.Up, 0);

        for (var row = 0; row < map.Count; row++)
        {
            for (var col = 0; col < map[row].Length; col++)
            {
                if (map[row][col] == '^')
                {
                    guard.Position = (row, col);
                    seenPositions.Add((row, col));
                    guard.Progress = 1;
                }
            }
        }

        for (;;)
        {
            try
            {
                var nextPosition = (row: 0, col: 0);
                switch (guard.Direction)
                {
                    case Direction.Up:
                        nextPosition = (row: guard.Position.Row - 1, col: guard.Position.Col);
                        break;
                    case Direction.Right:
                        nextPosition = (row: guard.Position.Row, col: guard.Position.Col + 1);
                        break;
                    case Direction.Down:
                        nextPosition = (row: guard.Position.Row + 1, col: guard.Position.Col);
                        break;
                    case Direction.Left:
                        nextPosition = (row: guard.Position.Row, col: guard.Position.Col - 1);
                        break;
                    default:
                        // hehe
                        _ = map[-1][-1];
                        break;
                }

                if (map[nextPosition.row][nextPosition.col] != '#')
                {
                    guard.Position = nextPosition;

                    if (!seenPositions.Contains(nextPosition))
                    {
                        guard.Progress++;
                    }

                    seenPositions.Add(nextPosition);
                }
                else
                {
                    guard.Direction = (Direction)(((int)guard.Direction + 1) % 4);
                }
            }
            catch (Exception e) when (e is IndexOutOfRangeException or ArgumentOutOfRangeException)
            {
                solution = guard.Progress;
                break;
            }
        }

        _output.WriteLine($"\tPart One: {solution}");
        Assert.Equal(5101, solution);
    }

    [Fact]
    public void PartTwo()
    {
        var solution = 0;
        var map = File.ReadAllLines(InputData).ToList();
        var seenPositions = new HashSet<(int Row, int Col, Direction Direction)>();
        var originalPosition = (0, 0);
        var originalDirection = Direction.Up;
        var guard = new Guard((0, 0), Direction.Up, 0);

        for (var row = 0; row < map.Count; row++)
        {
            for (var col = 0; col < map[row].Length; col++)
            {
                if (map[row][col] == '^')
                {
                    guard.Position = (row, col);
                    originalPosition = guard.Position;
                    seenPositions.Add((row, col, guard.Direction));
                    guard.Progress = 1;
                }
            }
        }

        for (var blockRow = 0; blockRow < map.Count; blockRow++)
        {
            for (var blockCol = 0; blockCol < map[blockRow].Length; blockCol++)
            {
                if (map[blockRow][blockCol] is '#' or '^')
                {
                    continue;
                }
                else
                {
                    seenPositions.Clear();
                    map[blockRow] = $"{map[blockRow][0..blockCol]}${map[blockRow][(blockCol+1)..]}";
                    guard.Position = originalPosition;
                    guard.Direction = originalDirection;
                    seenPositions.Add((guard.Position.Row, guard.Position.Col, guard.Direction));
                }

                for (;;)
                {
                    try
                    {
                        var nextPosition = (row: 0, col: 0);
                        switch (guard.Direction)
                        {
                            case Direction.Up:
                                nextPosition = (row: guard.Position.Row - 1, col: guard.Position.Col);
                                break;
                            case Direction.Right:
                                nextPosition = (row: guard.Position.Row, col: guard.Position.Col + 1);
                                break;
                            case Direction.Down:
                                nextPosition = (row: guard.Position.Row + 1, col: guard.Position.Col);
                                break;
                            case Direction.Left:
                                nextPosition = (row: guard.Position.Row, col: guard.Position.Col - 1);
                                break;
                            default:
                                // hehe
                                _ = map[-1][-1];
                                break;
                        }

                        if (map[nextPosition.row][nextPosition.col] is not ('#' or '$'))
                        {
                            guard.Position = nextPosition;

                            if (seenPositions.Contains((nextPosition.row, nextPosition.col, guard.Direction)))
                            {
                                solution++;
                                break;
                            }

                            seenPositions.Add((nextPosition.row, nextPosition.col, guard.Direction));
                        }
                        else
                        {
                            guard.Direction = (Direction)(((int)guard.Direction + 1) % 4);
                        }
                    }
                    catch (Exception e) when (e is IndexOutOfRangeException or ArgumentOutOfRangeException)
                    {
                        break;
                    }
                }

                map[blockRow] = map[blockRow].Replace("$", ".");
            }
        }

        _output.WriteLine($"\tPart Two: {solution}");
        Assert.Equal(1951, solution);
    }
}