#include "raylib.h"
#include "raymath.h"

const int screenWidth = 1000;
const int screenHeight = 1000;

typedef struct Card
{
    Model *model;
    Texture2D *texture;
    Vector3 *pos;
} Card;

int main(void)
{
    InitWindow(screenWidth, screenHeight, "raylib [models] example - models loading");
    Camera camera = {0};
    camera.position = (Vector3){0, 0, 150};
    camera.target = (Vector3){0, 0, 0};
    camera.up = (Vector3){0, 1, 0};
    camera.fovy = 45.0f;
    camera.projection = CAMERA_CUSTOM;

    Model modelJennie01 = LoadModel("assets/models/card-jennie-01.obj");
    Texture2D textureJennie01 = LoadTexture("assets/images/texture-jennie-01.png");
    SetMaterialTexture(&modelJennie01.materials[0], MATERIAL_MAP_DIFFUSE, textureJennie01);
    float yAngle = 0;
    Vector3 posJennie03 = (Vector3){-50, 0, 0};

    Card cardJennie01 = (Card){0};
    cardJennie01.model = &modelJennie01;
    cardJennie01.texture = &textureJennie01;
    cardJennie01.pos = &posJennie03;

    SetTargetFPS(60);
    while (!WindowShouldClose())
    {
        float dt = GetFrameTime();
        Matrix matrixRotateY = MatrixRotateY(yAngle * DEG2RAD);
        cardJennie01.model->transform = matrixRotateY;

        if (IsMouseButtonPressed(MOUSE_LEFT_BUTTON) || IsMouseButtonPressed(MOUSE_BUTTON_LEFT))
        {
            printf("mouse clicked");
            Vector2 mousePos = GetMousePosition();
            // Unproject the mouse position into a ray in 3D space
            Ray ray = GetMouseRay(mousePos, camera);
            BoundingBox modelBoundingBox = GetMeshBoundingBox(cardJennie01.model->meshes[0]);
            RayCollision collision = GetRayCollisionBox(ray, modelBoundingBox);
            printf(collision);
            if (collision.hit)
            {
                printf("hit!\n");
            }
        }

        BeginDrawing();
        ClearBackground(RAYWHITE);
        DrawRectangle(0, 0, screenWidth, screenHeight, BEIGE);
        BeginMode3D(camera);
        DrawModel(*cardJennie01.model, *cardJennie01.pos, 1, WHITE);

        DrawGrid(20, 10);
        EndMode3D();
        DrawFPS(10, 10);
        EndDrawing();
        yAngle += 50 * dt;
    }
    // {

    //                                                           // Calculate 3D model's bounding box in screen space
    //                                                           // bbox := rl.GetMeshBoundingBox(*cardJennie01.Model.Meshes)

    //                                                           if rl.IsMouseButtonPressed(rl.MouseLeftButton){
    //                                                               ray : = rl.GetMouseRay(rl.GetMousePosition(), camera)
    //                                                                       p :
    //                                                                   = ray.Position
    //                                                                         fmt.Println("ray.Position", p)
    //                                                                             fmt.Println("jennie.Position", cardJennie01.Position)
    //                                                                                 rl.DrawCube(ray.Position, 10, 10, 10, rl.Red)
    //                                                               // modelBoundingBox := rl.GetModelBoundingBox(*cardJennie01.Model)
    //                                                           }

    // }

    UnloadModel(modelJennie01);
    UnloadTexture(textureJennie01);
    CloseWindow();
}

// type Card struct
// {
// }

//     func
//     NewCard(model *rl.Model, texture *rl.Texture2D, position *rl.Vector3) *
//     Card
// {
//     return &Card
//     {
//     Model:
//         model,
//             Texture : texture,
//                       Position : position,
//     }
// }

// func(c *Card) Draw()
// {
// }
